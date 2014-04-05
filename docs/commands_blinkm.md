# Functions

## FirmwareVersion

Returs the version of the current Frimware.

### Returns

- **string** - Firmware Version 

#### API Command

**FirmwareVersionC**

## Color

Returns the color of the LED.

#### Returns

- **[]byte** - the RGB LED color

#### API Command

**ColorC**

## Rgb(r byte, g byte, b byte)

Sets the RGB color.

#### Params

- **r** - **byte** - Red
- **g** - **byte** - Green
- **b** - **byte** - Blue

#### API Command

**RgbC**

## Fade(r byte, g byte, b byte)

Fades the RGB color.

#### Params

- **r** - **byte** - Red
- **g** - **byte** - Green
- **b** - **byte** - Blue

#### API Command

**FadeC**